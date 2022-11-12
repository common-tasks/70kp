let vehicle = { wheels : 4 }; // object assigned to variable named vehicle
let car = { seats : 5 }; // object assigned to variable named car
let driver = {} // empty object assigned to variable named driver

// Print all objects and __proto__ property for each variable
console.log(`vehicle:`, vehicle, vehicle.__proto__);
console.log(`car:`, car, car.__proto__);
console.log(`driver:`, driver, driver.__proto__);