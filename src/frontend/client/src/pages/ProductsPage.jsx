import ProductItem from "../components/ProductItem";

import { Flex } from "@chakra-ui/react";

const ProductsPage = () => {
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
      <ProductItem />
      <ProductItem />
      <ProductItem />
      <ProductItem />
      <ProductItem />
    </Flex>
  );
};

export default ProductsPage;
