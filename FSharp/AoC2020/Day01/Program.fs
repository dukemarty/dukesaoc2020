// Learn more about F# at http://fsharp.org

open System
open System.IO


// All ordered picks {x_i1, x_i2, .. , x_ik} of k out of n elements {x_1,..,x_n}
// where i1 < i2 < .. < ik
let picks n L = 
    let rec aux nleft acc L = seq {
        match nleft,L with
        | 0,_ -> yield acc
        | _,[] -> ()
        | nleft,h::t -> yield! aux (nleft-1) (h::acc) t
                        yield! aux nleft acc t }
    aux n [] L


let readLines (filePath:string) = seq {
    use sr = new StreamReader (filePath)
    while not sr.EndOfStream do
        yield sr.ReadLine ()
}

let solveForN (data:seq<int>) n:int =
    let r = picks n (List.ofSeq data) |> Seq.filter (fun l -> List.sum l = 2020)
    let res = List.fold (fun acc i -> acc * i) 1 (Seq.head r)
    res

let solvePart1 (data:seq<int>) = 
    printfn "Part 1: Two numbers who add up to 2020, their product\n-----------------------------------------------------"
    printfn $"  Result: {solveForN data 2}"


let solvePart2 (data:seq<int>) =
    printfn "Part 2: Three numbers who add up to 2020, their product\n-------------------------------------------------------"
    printfn $"  Result: {solveForN data 3}"

[<EntryPoint>]
let main argv =
    printfn "Day 01: Report Repair\n====================="
    let data = readLines "RawData.txt" |> Seq.map int
    solvePart1 data
    solvePart2 data
    0 // return an integer exit code
