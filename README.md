# hyst

Write a histogram in Terminal.

[![Build Status](https://travis-ci.org/winebarrel/hyst.svg?branch=master)](https://travis-ci.org/winebarrel/hyst)

## Installation

```
brew install https://raw.githubusercontent.com/winebarrel/hyst/master/homebrew/hyst.rb
```

## Usage

```
Usage of hyst:
  -w int
      Width
```

```
$ ruby -e 'puts 1000.times.map { %w(foo bar zoo)[rand(3)] }' > rand.txt

$ head rand.txt
zoo
zoo
bar
foo
zoo
bar
zoo
zoo
bar
zoo

$ cat rand.txt | hyst
zoo  351  ##################################################
foo  341  ################################################
bar  308  ###########################################
```

```
$ ruby -rdistribution -e 'normal = Distribution::Normal.rng(1); puts 1_000.times.map {normal.call * 100}' > gauss.txt

$ head gauss.txt
213.57863028255827
-103.40921634097238
-11.296606231306994
162.5407288364103
110.5574029438751
24.64428067386242
30.02832601441755
100.63481149177946
-23.661175314447913
185.2196657484519

$ cat gauss.txt | hyst -w 50
-250    1
-200    2
-150    4
-100   12  ##
 -50   51  ##########
   0  247  ##################################################
  50  198  ########################################
 100  198  ########################################
 150  137  ###########################
 200   96  ###################
 250   36  #######
 300   15  ###
 350    2
 400    1
```

## Installation

### OS X

```sh
brew install https://raw.githubusercontent.com/winebarrel/hyst/master/homebrew/hyst.rb
```
