import { useParams } from "react-router-dom";
import products from "../data";
const SingleProduct = () => {
  const { productId } = useParams();
  const product = products.find((product)=>product.id===productId);
  const {image,name}=product;
  return (
    <section className="section product">
      <img src={image} alt={name}></img>
      <h5>{name}</h5>
    </section>
  );
};

export default SingleProduct;
