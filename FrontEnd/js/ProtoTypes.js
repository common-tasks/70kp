var arr1 = [1,2,3,4];
var arr2= new Array();
console.log(Array.prototype);
Array.prototype.lol ='ha ha ha';
Array.prototype.loool='huhuhu';
console.log(Array.prototype);
console.log(Array.lol);
console.log(arr1.lol);
console.log(arr1);

arr1.push(99)

console.log(arr1);

arr1.pop();

console.log(arr1);

arr1.shift();
console.log(arr1);
arr1.unshift(1001);
console.log(arr1);

var arr3 = [23,34,45,56,67];
var sarr3= arr3.toString();
console.log(arr3,sarr3);
console.log(typeof arr3, typeof sarr3);