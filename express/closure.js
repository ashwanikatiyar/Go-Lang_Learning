
function increament(value){
    value++
    return value
}


a = 10
increament(a)
console.log(a) // still 10 (no change to the value)

// Now let make a Closure function (From go)

function change(startValue){
    let count = startValue
    return function(){
        count++
        return count
    }
}

a1 = change(0)
a2 = change(0)

console.log(a1()) // 1 
console.log(a1()) // 2
console.log(a2()) //Twist , a2 is independent so starts with 0
console.log(a1()) // 3 , a1 still continues  