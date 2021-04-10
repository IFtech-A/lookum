import React from 'react'
import ProductDetailed from '../components/ProductDetailed'
import { useParams } from 'react-router-dom'

function ProductDetailedScreen() {

    let { id } = useParams();
    console.log("screen" + id);

    return (
        <ProductDetailed id={id} />
    )
}

export default ProductDetailedScreen;