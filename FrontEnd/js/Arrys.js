let arr1 =["hello","world","I","am","here"];
console.log(arr1);
let arr2 = new Array();
arr2.push("ji");
arr2.push('haan');
console.log(arr2);
arr2.unshift("mai aaya");
console.log(arr2);

arr2.splice(arr2.length-1,1);
console.log(arr2);

arr2.push("haan");

const arr3 = arr2.slice(0,arr2.length);

console.log(arr2);
console.log(arr3.shift());