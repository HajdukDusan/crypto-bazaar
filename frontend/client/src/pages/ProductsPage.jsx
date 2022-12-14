import { Flex } from "@chakra-ui/react";
import { useEffect } from "react";
import { useState } from "react";
import Item from "../components/Item";

const ProductsPage = () => {
  const [items, setItems] = useState([]);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await fetch("http://localhost:9001/book");
      const json = await response.json();
      setItems(json.rows);
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <Flex
      justifyContent="center"
      w={["90vw", "90vw", "100vw"]}
      direction={["column", "column", "row", "row"]}
      justify="center"
      alignItems="center"
      mx="auto"
      maxW="10xl"
    >
      {items.map((item) => (
        <Item item={item} />
      ))}
    </Flex>
  );
};

export default ProductsPage;
