var obj1 = { name1: "anurag", name2: "draupadi murmu" };
// console.log(obj1);

console.log(Object.keys(obj1));
console.log(Object.values(obj1));
console.log(Object.entries(obj1));

var keys = Object.keys(obj1);
for (let index = 0; index < keys.length; index++) {
  console.log(keys[index], obj1[keys[index]]);
}

for (var i in obj1) {
  console.log(i, obj1[i]);
}
