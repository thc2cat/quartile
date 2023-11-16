# Using quartile to determine abnormal values

![go.yml](https://github.com/thc2cat/quartile/actions/workflows/go.yml/badge.svg)

This tool use quartile method to detect abnormal values from a dataset using various méthods.

## Typical usage

```shell
 grep pattern |awk '{ print $fields }'\
  | sort| uniq -c | sort -rn | head -30 \
  | quartile [options]
```

optionals calc are :

* -m minimal value : to reduce low number parasites

* Default is quartile calc
* -f deviance Tuckey Factor : if normal quartile 1.5 is insufficient
* -M for 3xAvg deviance
* -B for Boxplot méthod
* -Z for Z-Score méthod (using avg)
* -D for Z-Score modified method ( using median point )

* Quiet mode only exit code on abnormal values, and no outputs
* Print quartiles values show Median, Q1, Q2
* Print limits show limits used in quartile methods

Output of deviant values are prefixed with ">" or "<".

>> `Memory Warning :`
>> data file is fully loaded in memory before sorting and printing.
>> But usually you only search limited values.

## Exemple

What log file is suspect for you ?

```shell
# wc -l 2021-03*/mx.domain/mail.log
   178134 2021-03-01/mx.domain/mail.log
   184757 2021-03-02/mx.domain/mail.log
   198314 2021-03-03/mx.domain/mail.log
   211160 2021-03-04/mx.domain/mail.log
   168873 2021-03-05/mx.domain/mail.log
   101494 2021-03-06/mx.domain/mail.log
    97250 2021-03-07/mx.domain/mail.log
   236236 2021-03-08/mx.domain/mail.log
   253316 2021-03-09/mx.domain/mail.log
   181050 2021-03-10/mx.domain/mail.log
   172090 2021-03-11/mx.domain/mail.log
   165354 2021-03-12/mx.domain/mail.log
    97664 2021-03-13/mx.domain/mail.log
   110994 2021-03-14/mx.domain/mail.log
   193601 2021-03-15/mx.domain/mail.log
   210465 2021-03-16/mx.domain/mail.log
   159333 2021-03-17/mx.domain/mail.log
   166172 2021-03-18/mx.domain/mail.log
   173893 2021-03-19/mx.domain/mail.log
    89610 2021-03-20/mx.domain/mail.log
    82982 2021-03-21/mx.domain/mail.log
   290674 2021-03-22/mx.domain/mail.log
   188349 2021-03-23/mx.domain/mail.log
    51808 2021-03-24/mx.domain/mail.log
  3963573 total                                              

$ wc -l 2021-03*/mx.domain/mail.log | rg mail.log | quartile

output nothing, cause we don t have a really big deviation, 
but with -B (Box Method), output low and high values

$ go build ; ./quartile.exe -B < data2
< 89610 2021-03-20/mx.domain/mail.log
< 82982 2021-03-21/mx.domain/mail.log
> 290674 2021-03-22/mx.domain/mail.log
< 51808 2021-03-24/mx.domain/mail.log

When dispersion is low, you can also try 
lowering Tukey factor ( 1.5 default ) 

$ go build ; ./quartile.exe -f .8 < data2
> 290674 2021-03-22/mx.domain/mail.log

When there is spikes in your data , Z-Score method 
is usually effective (exemple spammers)

$ go build ; ./quartile.exe -Z < data
> 881 thomas.Banslish@vsuq.org

```
