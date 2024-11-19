module Main where

main = do
    let list = [386, 19, 2024]

    let multiplicatedList = map (*3) list

    let reverseList = reverse multiplicatedList

    print (head reverseList)
