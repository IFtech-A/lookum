
const API_SERVER = "https://lookum.org/"

export const fetchProducts = async () => {
  const response = await fetch(API_SERVER + "products")
  if (!response.ok) {
    throw new Error(response.statusText);
  }
  const productsJson = await response.json()
  return productsJson
}

export const fetchProduct = async (id) => {
  const response = await fetch(API_SERVER + "product/" + id)
  if (!response.ok) {
      console.log(response)
    throw new Error(response.statusText);
  }
  const productJson = await response.json()
  return productJson
}

export const createProduct = async (product) => {
  
  return 4
}

export const uploadProductImages = async (product) => {
  return true
}