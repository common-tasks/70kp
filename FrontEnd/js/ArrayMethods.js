const arr = [12, 2, 34, 67, 90, 88, 2, 3, 4, 8];

arr.forEach((val, i, array) => {
  console.log(val, i, array);
});

const arr2 = arr.filter((x) => x > 80);
console.log(arr2);
console.log(arr);

const arr3 = arr.find((x) => x > 5);
console.log(arr3, typeof arr3);

const arr4 = arr.map((x) => x * 2);
console.log(arr4);

const sum = arr.reduce((prev, curr) => prev + curr);
console.log(sum);

let total = 0;
for (let index = 0; index < arr.length; index++) {
  total = total + arr[index];
}
console.log(total);

var mixarr= ['hello',1,2,3,4,NaN,Boolean,'world'];

var strArr= mixarr.map((x)=>typeof x ==='string'?1:0);
console.log(strArr);
var numberofstr= strArr.reduce((prev,curr)=>prev+curr)
console.log(numberofstr);

var twodarr= [[1],[2,3,4],[3,4]];
var sarr= twodarr.map((x)=>x.length);
console.log(sarr);
var sumsarr= sarr.reduce((prev,curr)=>prev+curr);
console.log(sumsarr);

const greetMe = ()=>{
    console.log('hello me');
}
export {greetMe};