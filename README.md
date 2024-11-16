
# Week 8 Assignment - Modern Applied Statistics with Go

This repository is a comparison of statistical analysis between Go and R. In this repository, there is a both a Go and R version of a program attempting to find the sample median of a set of C02 concentractions. The R version of this program seems to report a standard deviation while I designed my program to report standard error. I say this because the standard deviation reported by the R boot function, ~1.3, was around the value of standard deviation that I calculated with Go. When converting that value with Go to standard error by dividing by the square root of the size of the sample, I got my eventual standard error of ~0.04. The boot function report's its value with the name standard error, so I am not sure why it reports a standard deviation. Either way, this represents the main difference between the two functions. 

## How to use this tool and examples:
* After cloning this git repository you can `go run main.go
* You will receive the Go estimations of sample median and standard error. You can also run the test in the repo to make sure standard error is being calculated correctly and benchmark the program with the benchmark test in the repo. 

In my tests, R was actually faster by a large margin with around a 0.15 second proccessing time to around a 1 second processing time for Go. This might be a sign that I wrote slow running Go code, but that is another factor in the favor of R. The solution to this problem in R was several easy to write and research lines. The solution in Go was around a 100 very hard to write lines. Given the drop off in performance and the extreme difference in ease of use, I think it makes sense for this task to still be undertaken in R.

