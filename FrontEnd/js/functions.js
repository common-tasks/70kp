import {greetMe} from './ArrayMethods';
function func1(msg) {
  console.log(msg);
  greetMe();
}

const func2 = function (arg) {
  console.log(arg);
};

const func3 = (arg) => {
  console.log(arg);
};

const func4 = (func3, arg) => {
  func3(arg);
};
const func5 = function () {
  console.log("i will see you later");
};

func1("hello");
func2("world");
func3("halo");
func4(func3, "hola");
func5();
