module Main where

factorial :: Int -> Int
factorial 0 = 1
factorial n = n * factorial (n - 1)

main = do
    let number = 4
    let result = if number > 0
                 then factorial number
                 else number * 2
    print result