import {
  Button,
  Flex,
  Box,
  Text,
  useColorMode,
  Center,
} from "@chakra-ui/react";
import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <>
      <Box maxW="2xl" mx={"auto"} pt={1} px={{ base: 2, sm: 12, md: 17 }}>
        <Center>
          <Text fontSize="6xl">Crypto Bazaar</Text>
        </Center>
      </Box>
      <Flex
        flexDirection="row"
        justifyContent="space-between"
        alignItems="center"
        width="100%"
        as="nav"
        p={4}
        mx="auto"
        maxWidth="1150px"
        marginBottom="40px"
        maxW="8xl"
      >
        <Box>
          <Link to="/">
            <Button
              fontWeight={["medium", "medium", "medium"]}
              fontSize={["xs", "sm", "lg", "xl"]}
              variant="ghost"
              _hover={{ bg: "rgba(0,0,0,.2)" }}
              padding="1"
              color="white"
              letterSpacing="0.65px"
            >
              Ethereum Transaction Crawler
            </Button>
          </Link>
        </Box>

        <Box>
          <Link to="/historical-value">
            <Button
              fontWeight={["medium", "medium", "medium"]}
              fontSize={["xs", "sm", "lg", "xl"]}
              variant="ghost"
              p={[1, 4]}
            >
              Historically Available Value
            </Button>
          </Link>
        </Box>
      </Flex>
    </>
  );
};

export default Navbar;
