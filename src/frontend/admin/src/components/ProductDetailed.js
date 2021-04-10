import React from 'react';

import { Spinner } from 'react-activity'

import { fetchProduct } from './api';

const styles = {
    bodyContainer: {
        display: "flex",
        justifyContent: 'center',
        backgroundColor: "#ccc"
        // padding: 10,
    },
    container: {
        margin: 10,
        marginHorizontal: 15,
        elevation: 3,
        shadowColor: "#000",
        // shadowOffset: { width: 2, height: 2 },
        shadowOpacity: 0.5,
        shadowRadius: 2
    },
    errContainer: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
    },
    imageContainer: {
        elevation: 2,
        margin: 0,
    },
    productHeader: {
        paddingRight: 10,
        paddingLeft: 10,
        paddingBottom: 10,
        backgroundColor: 'white',
        borderBottomLeftRadius: 10,
        borderBottomRightRadius: 10,
        marginBottom: 15,
    },
    title: {
        margin: 0,
        fontSize: 26,
        fontStyle: 'italic',
    },
    priceTag: {
        margin: 0,
        fontSize: 24,
        color: 'red',
    },
    description: {
        margin: 0,
        fontSize: 18,
        color: '#3d3d3d',
    },
};

function _ProductDetails(props) {



    return (
        <div style={styles.bodyContainer}>
            <div style={styles.container}>
                <div style={styles.imageContainer}>
                    {/* Header */}
                    <img
                        alt=""
                        style={{
                            height: 500,
                            borderTopLeftRadius: 10,
                            borderTopRightRadius: 10,
                        }}
                        src={props.images[0].file_uri}
                    />
                </div>
                {/* Body */}
                <div style={styles.productHeader}>
                    <p style={styles.title}>{props.title}</p>
                    <p style={styles.priceTag}>${props.price}.00</p>
                    <p style={styles.description}>{props.desc}</p>
                </div>

                {/* Footer */}

            </div>
        </div>
    );
}

export default class ProductDetails extends React.Component {

    state = {
        isReady: false,
        // imageSize: null,
    };

    constructor(props) {
        super(props);
        console.log(props);
    }

    _getProduct = async () => {
        try {
            const product = await fetchProduct(this.props.id);
            this.setState({ ...product, isReady: true });
            //   await this._getImageSizes(product.images);
        } catch (err) {
            console.error('fetching product failed: ' + err.message);
        }
    };

    //   _getImageSizes = async (images) => {
    //     const imageSizes = [];
    //     for (let i = 0; i < images.length; i++) {
    //       let imageSize = {};
    //       await Image.getSize(images[i].file_uri, (w, h) => {
    //         console.log(w, h);
    //         imageSize = { width: w, height: h };
    //       });
    //       imageSizes.push(imageSize);
    //     }
    //     this.setState({ imageSizes });
    //   };

    componentDidMount() {

        this._getProduct();
    }

    render() {
        return this.state.isReady ? (
            <_ProductDetails {...this.state} />
        ) : (
            <div style={styles.errContainer}>
                <Spinner />
            </div>
        );
    }
}
