package main

import (
	"math"
)

// Exemple de calcul
// 10, 25, 30, 40, 41, 42, 50, 55, 70, 101, 110, 111 => 12 Valeurs
//
//	40      42+50       70
//
// => M = 42+50/2 => 46
// => Q1 = 12/4=> 4 eme valeur => 40
// => Q2 = 12*3/4 => 9 eme valeur => 70
//
// Dans la fonction qui suit, R[0]==Mediane, R[1]=Q1 et R[2]=Q3
//
// https://fastercapital.com/fr/contenu/Valeurs-aberrantes-dans-les-quartiles---identification-des-valeurs-extremes-dans-l-ensemble-de-donnees.html#Introduction-aux-valeurs-aberrantes-et-aux-quartiles

// quartileCalcf32 return quartiles Q1-3 from []float32 array
func quartileCalcf32(data []float32) [3]float32 { // Return Q1 Q2 Q3
	var Q [3]float32
	var n float64
	var i int

	sortSlice(data)

	N := len(data)
	if N < 4 {
		// log.Printf("bad quartileCalc %v", d)
		return [3]float32{-1, 0, 0}
	}

	// Calcul de la mediane Q[0]
	if N%2 == 0 { // Pair
		Q[0] = (data[N/2-1] + data[N/2]) / 2
	} else {
		Q[0] = data[(int)(math.Ceil((float64)(N)/2))-1]
	}

	// Calc Q1
	n = (float64)(N) / 4
	if isDecimal(n) { // entier
		i = (N / 4) + 1
	} else {
		i = (int)(math.Ceil(n))
	}
	Q[1] = data[i-1]

	// calc Q3
	n = (float64)(N) * 3 / 4
	if isDecimal(n) { // entier
		i = (N * 3) / 4
	} else {
		i = (int)(math.Ceil((float64)(N) * 3 / 4))
	}
	Q[2] = (float32)(data[i-1])

	return Q
}

// isDecimal function test if float is decimal
func isDecimal(n float64) bool {
	return math.Floor(n) == n
}

/*
 *
 1. Méthode de la boîte à moustaches

 La méthode Boxplot est une méthode graphique permettant de détecter les valeurs
 aberrantes dans les données. Un boxplot affiche la distribution d'un ensemble de
 données en affichant la médiane, les quartiles et la plage des données. La
 fourchette est définie comme la différence entre les quartiles supérieur et
 inférieur. Tout point de données qui se situe en dehors des moustaches du
 boxplot est considéré comme une valeur aberrante.
 La méthode du boxplot est simple, rapide et efficace.

2. Méthode Z-Score.

 La méthode Z-score est une méthode statistique permettant de détecter les
 valeurs aberrantes dans les données. Le score Z mesure la distance entre un
 point de données et la moyenne de l'ensemble de données en termes d'écart type.
 Tout point de données qui s’écarte de plus de trois écarts types de la moyenne
 est considéré comme une valeur aberrante. La méthode du score Z est efficace
 mais peut être influencée par les valeurs extrêmes de l'ensemble de données.

3. Méthode Z-Score modifiée

 La méthode Z-score modifiée est une variante de la méthode Z-score. Il est plus
 robuste et moins influencé par les valeurs extrêmes de l’ensemble de données.
 Le score Z modifié mesure la distance entre un point de données et la médiane
 de l'ensemble de données en termes d'écart absolu médian (MAD). Tout point de
 données éloigné de plus de trois MAD de la médiane est considéré comme une
 valeur aberrante. La méthode Z-score modifié est efficace et robuste.

*/

// ZScoreCalF32 calculate median and ecart type from []entries > mini value
func ZScoreCalF32(data []entries, mini float64) (float32, float32) {
	//
	// ZScore => Calculer la moyenne, l'ecart type
	// si l'ecart type > 3x l'ecart moyen alors suspect
	//
	var somme, sommeecarts float32

	var count int

	for _, v := range data {
		if (mini == 0) || (v.value >= float32(mini)) {
			somme += v.value
			count++
		}
	}
	ecarts := make([]float32, count)
	moyenne := somme / (float32)(count)

	var i int
	for _, v := range data {
		if (mini == 0) || (v.value >= float32(mini)) {
			ecarts[i] = myAbs(v.value - moyenne)
			sommeecarts += ecarts[i]
			i++
		}
	}

	ecartmoyen := sommeecarts / (float32)(count)

	return moyenne, ecartmoyen
}

func ZScoreMADCalF32(data []entries, mini float64) (float32, float32) {
	//
	// ZScore => Calculer la moyenne, l'ecart type
	// si l'ecart type > 3x l'ecart moyen alors suspect
	//
	var sommeecarts float32

	var count, myidx int

	// Par rapport au nombre, trouver le point médian (!! pas la moyenne)
	Arr := getNumbers(data, mini)
	sortSlice(Arr)

	N := len(Arr)
	if N == 0 {
		return 0, 0
	}

	if N%2 == 1 { // imPair
		myidx = (N + 1) / 2
	} else {
		myidx = N / 2
	}

	MAD := Arr[myidx] // Median Absolute Data point

	for _, v := range data {
		if (mini == 0) || (v.value >= float32(mini)) {
			count++
		}
	}
	ecarts := make([]float32, count)

	var i int
	for _, v := range data {
		if (mini == 0) || (v.value >= float32(mini)) {
			ecarts[i] = myAbs(v.value - MAD)
			sommeecarts += ecarts[i]
			i++
		}
	}

	ecartmoyen := sommeecarts / (float32)(count)

	return MAD, ecartmoyen
}

func myAbs(x float32) float32 {
	switch {
	case x < 0:
		return -x
	}
	return x
}
