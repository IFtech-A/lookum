import React from 'react';

import { Spinner } from 'react-activity'
import { fetchProducts } from '../components/api';
import Product from './Product';



export default class ProductsList extends React.Component {
  state = {
    products: [],
    err: null,
    isReady: false,
  };


  _fetchProducts = async () => {
    try {
      const products = await fetchProducts();
      this.setState({
        products: products,
        err: null,
        isReady: true,
      });
    } catch (err) {
      console.error('fetching products failed: ' + err.Message);
      this.setState({
        err: 'fetching products failed: ' + err.Message,
        isReady: true,
      });
    }
  };

  showProduct = (id, title) => {
    this.props.navigation && this.props.navigation.push('ProductDetails', { id, productTitle: title });
    
  };

  componentDidMount() {
    this._fetchProducts();
  }

  renderProduct = (item) => (
    <Product key={item.id} {...item} />
  );

  render() {
    if (this.state.err) {
      return (
        <div
          style={styles.errContainer}>
          {this.state.err}
        </div>
      );
    }
    return !this.state.isReady ? (
      <div style={styles.errContainer}>
        <Spinner />
      </div>
    ) : (
      <div
        style={styles.bodyContainer}>
        {this.state.products.map(this.renderProduct)}
      </div>
    );
  }
}

const styles = {
  bodyContainer : {
    display: "flex",
    justifyContent: "center",
  },
  errContainer: {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
  }
}