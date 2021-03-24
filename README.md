# Using quartile to determine abnormal values

This tool use quartile method to print out abnormal values from a dataset using quartile method.

optionals calc are :

* Median output : a value is abnormal if 3x>Median value
* minimal value : to reduce low number parasites
* devianceFactor : if normal quartile 1.5 is insufficient

* Quiet mode only exit code on abnormal values, and no outputs
* Print quartiles values show Median, Q1, Q2

> `Memory Warning :`
> data is fully loaded in memory before sorting and printing.

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

# wc -l 2021-03*/mx.domain/mail.log | awk '!/total/{ print $1 }' | quartile

output nothing, cause we don t have a really big deviation

Using Q1 values from "-p" option , we set minimal value

# wc -l 2021-03*/mx.domain/mail.log | awk '!/total/{ print $1 }' | quartile -p -m 110994
== Q1=172090 Mediane=184757 Q3=210465 ==
290674

So we should have a look to 2021-03-22/mx.domain/mail.log which should be your first suspect.
```
