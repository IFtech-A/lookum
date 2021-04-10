import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import Products from './components/Products'
import ProductDetailedScreen from './screens/ProductDetailedScreen'
import { Nav, Navbar } from 'react-bootstrap'

import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom'

import 'bootstrap/dist/css/bootstrap.min.css';
import CreateProductScreen from './screens/ProductCreateScreen'


export default class App extends React.Component {

    render() {
        return (
            <Router>
                <Navbar bg="dark" variant="dark">
                    <Navbar.Brand href="/">Lookum</Navbar.Brand>
                    <Nav>
                        <Nav.Link href="/createProduct">Create Product</Nav.Link>
                        <Nav.Link href="/">Products</Nav.Link>
                    </Nav>
                </Navbar>
                <Switch>

                    <Route exact path="/">
                        <Products />
                    </Route>
                    <Route path="/createProduct" children={<CreateProductScreen />} />
                    <Route path="/product/:id" children={<ProductDetailedScreen />} />
                    <Route path="/">
                        <Products />
                    </Route>
                </Switch>
            </Router>
        )
    }
}

ReactDOM.render(
    <App />,
    document.getElementById('root')
)