var names =["draupadi","murmu","james","khalkho"]
var dnames= names.slice(0,10);
names[0]="mahamahim draupadi"
console.log(dnames);
console.log(names);
names.splice(0,0,"chaplus");
console.log(names);
console.log(dnames);

var emp =[["anurag",110],["gama",100000],["bhola",200000]]
console.log(emp);
var emp4=["babua",500]
emp.push(emp4);
console.log(emp);
emp[3]=["rama",1];
console.log(emp);
emp.splice(3,1);
console.log(emp);
var e2= emp.slice();
console.log(e2.length);

names.forEach(element => {
    console.log(element);
});