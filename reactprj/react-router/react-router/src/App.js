import { BrowserRouter, Route, Routes } from "react-router-dom";
import Home from "./pages/Home";
import Products from "./pages/Products";
import About from "./pages/About";
import Error from "./pages/Error";

function App() {
  return (
    <BrowserRouter>
      <nav>Navigation Bar</nav>
      <Routes>
        <Route path="/" element={<Home />}>
          <Route path="products" element={<Products />}></Route>
          <Route path="home" element={<Home />}></Route>
          <Route path="about" element={<About />}></Route>
          <Route path="*" element={<Error />}></Route>
        </Route>
      </Routes>
      <footer>Copyright @Anurag Kumar</footer>
    </BrowserRouter>
  );
}

export default App;
