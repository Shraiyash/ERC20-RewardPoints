import { useState } from "react";
import api from "../utils/api";

export default function ProductPage() {
  const [productValue, setProductValue] = useState("");
  const [response, setResponse] = useState(null);

  const handlePurchase = async () => {
    try {
      const res = await api.buyProduct(productValue);
      setResponse(res.data);
    } catch (error) {
      setResponse({ error: error.message });
    }
  };

  return (
    <div>
      <h1>Buy Product</h1>
      <input
        type="number"
        placeholder="Enter product value"
        value={productValue}
        onChange={(e) => setProductValue(e.target.value)}
      />
      <button onClick={handlePurchase}>Buy Now</button>
      {response && (
        <div>
          {response.error ? (
            <p>Error: {response.error}</p>
          ) : (
            <p>Success: {JSON.stringify(response)}</p>
          )}
        </div>
      )}
    </div>
  );
}