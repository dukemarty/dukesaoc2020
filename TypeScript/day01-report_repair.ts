import { readFileSync } from 'fs';
import { exit } from 'process';


function getElementsForSum(numbers: number[], count: number, targetSum: number): number[] {
    if (count == 2){
        // console.log(`Reached final case, only 2 number required to reach ${targetSum} from array of length ${numbers.length}`)
        for (const { i, v } of numbers.map((v, i) => ({ i, v }))){
            for (var j=i+1;+j < numbers.length; ++j){
                if (v + numbers[j] == targetSum){
                    return [ v, numbers[j] ]
                }
            }
        }
    } else {
        for (let _i in numbers){
            if (targetSum > numbers[_i] + count - 1){
                let res = getElementsForSum(numbers.slice(+_i+1, numbers.length), count-1, targetSum - numbers[+_i])
                if (res != null){
                    return [numbers[+_i]].concat(res)
                }
            }
        }
    }

    return null
}


console.log("Day 01: Report Repair\n=====================")

const lines = readFileSync('./day01-rawdata.txt', 'utf-8').split("\n")
console.log("Numbers from input:")
console.log(lines.map(function (s, _){ return +s }))

console.log("Part 1: Two numbers who add up to 2020, their product\n-----------------------------------------------------")
let resPart1 = getElementsForSum(lines.map(function (s, _){ return +s }), 2, 2020)
console.log(`Found both numbers: ${resPart1[0]} + ${resPart1[1]} = ${resPart1[0]+resPart1[1]}`)
console.log("Multiplication result: %d\n", resPart1.reduce(function (acc, val, ind){ return acc * val }, 1))

console.log("Part 2: Three numbers who add up to 2020, their product\n-------------------------------------------------------")
let resPart2 = getElementsForSum(lines.map(function (s, _){ return +s }), 3, 2020)
console.log(`Found all three numbers: ${resPart2[0]} + ${resPart2[1]} + ${resPart2[2]} = ${resPart2[0]+resPart2[1]+resPart2[2]}`)
console.log("Multiplication result: %d\n", resPart2.reduce(function (acc, val, ind){ return acc * val }, 1))






