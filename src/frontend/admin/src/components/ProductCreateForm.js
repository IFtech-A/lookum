import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { createProduct, uploadProductImages } from '../components/api'

export default function CreateForm(props) {
    const { register, handleSubmit } = useForm();

    const onProductCreate = async (data) => {
        console.log("onProductCreate: " + JSON.stringify(data))
        const id = await createProduct(data)
    }

    return (
        <form onSubmit={handleSubmit(onProductCreate)}>
            <input type="hidden" value="1" {...register('vendor_id')} />
            <div>
                <label>Title</label>
                <input
                    type="text"
                    {...register('title', { required: true, maxLength: 80 })}
                />
            </div>
            <div>
                <label>SKU</label>
                <input
                    type="text"
                    {...register('sku', { required: true, maxLength: 20 })}
                />
            </div>
            <div>
                <label>Price</label>
                <input
                    type="number"
                    defaultValue="0"
                    placeholder="0"
                    {...register('price', {
                        required: true,
                        pattern: /^([1-9]{1}[0-9]*\.?[0-9]*)$|^([0]\.?[0-9]*)$/,
                    })}
                />
            </div>

            <div>
                <label>Discount</label>
                <input
                    type="number"
                    defaultValue="0"
                    placeholder="0"
                    {...register('discount', {
                        pattern: /^([1-9]{1}[0-9]*\.?[0-9]*)$|^([0]\.?[0-9]*)$/,
                    })}
                />
            </div>
            <div>
                <label>Quantity</label>
                <input
                    type="number"
                    defaultValue="0"
                    placeholder="0"
                    {...register('quantity', {
                        pattern: /^([1-9]{1}[0-9]*)$/,
                    })}
                />
            </div>
            <div>
                <label>Available for shop</label>
                <input
                    type="checkbox"
                    defaultChecked
                    {...register('shop_available')}
                />
            </div>
            <div>
                <label>Category</label>
                <select {...register('category', { required: true })}>
                    <option value="1">Top</option>
                    <option value="2">Bottom</option>
                </select>
            </div>

            <div>
                <label>Images</label>
                <input type="file" multiple {...register('images.0')} />
            </div>

                <input type="submit" />
        </form>
    );
}