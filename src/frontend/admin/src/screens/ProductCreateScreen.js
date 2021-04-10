import React from 'react'
import CreateProductForm from '../components/ProductCreateForm'

function CreateProductScreen() {

    const onSubmit = (data) => {
        alert(JSON.stringify(data))
    }

    return (
        <div style={{display:"flex", justifyContent:"center", alignItems:"center"}}>
            <CreateProductForm onSubmit={onSubmit} />            
        </div>
    )
}

export default CreateProductScreen
