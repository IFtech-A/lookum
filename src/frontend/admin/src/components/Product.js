import React from 'react';
import { Card } from 'react-bootstrap'

export const ProductHeader = (props) => {
    return (
        <div style={{ }}>
            <p style={styles.title}>{props.title}</p>
            <p style={styles.priceTag}>${props.price}.00</p>
        </div>
    );
};

export default class Product extends React.Component {

    _handleClick = () => {
        window.location.href = '/product/' + this.props.id
    }

    render() {
        return (
            <div
                style={styles.bodyContainer}
                onClick={this._handleClick}>
                <div>
                    <Card >
                        <Card.Img
                            style={styles.image}
                            src={this.props.images[0].file_uri}
                        />
                        <Card.Body>
                            <Card.Title>
                                {this.props.title}
                            </Card.Title>
                            <Card.Text>
                                ${this.props.price}.00
                            </Card.Text>

                        </Card.Body>
                    </Card>
                </div>
            </div>
        );
    }
}

const styles = {
    bodyContainer: {
        marginTop: 10,
        width: 400,
        display: "flex",
        flexDirection: "column",
        margin: 10,
        borderRadius: 10,
    },
    image: {
        height: 250,
    },
    title: {
        fontSize: 20,
        fontStyle: 'bold',
    },
    priceTag: {
        fontSize: 16,
        color: 'black',
    },
};
