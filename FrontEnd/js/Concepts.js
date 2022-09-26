console.log('hello world');

var fullMoon = true;

// Initialize a global variable
let species = "human";

if (fullMoon) {
  // Initialize a block-scoped variable
  let species = "werewolf";
  console.log(`It is a full moon. Lupin is currently a ${species}.`);
  var fullMoon = false;
  console.log(`is it a full moon ${fullMoon}`);
}

console.log(`It is not a full moon. Lupin is currently a ${species}.`);
console.log(`is it a full moon global : ${fullMoon}`);

// The code we wrote
// console.log(x);
// var x = 100;

// How JavaScript interpreted it
// var x;
// console.log(x);
// x = 100;

// Initialize x in the global scope
var x = 100;

function hoist() {
  // A condition that should not affect the outcome of the code
  if (false) {
    var x = 200;
  }
  console.log(x);
}

hoist();