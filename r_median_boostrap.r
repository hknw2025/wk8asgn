library(boot)

# data is originally from here
# data <- datasets::co2
# added read to make the benchmark test equivalent

t1 <- Sys.time()

data <- read.csv("/Users/harri/OneDrive/Documents/wk8asgn/data/co2_data.csv")$co2_data

b <- boot(data,statistic = function(x,i) median(x[i]),R = 1000)

median(b$t)

t2 <- Sys.time()
t2 - t1

#write data for go analysis
write.csv(data.frame(co2_data = data), "/Users/harri/OneDrive/Documents/wk8asgn/data/co2_data.csv")

