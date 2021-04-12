import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import Products from './components/Products'
import ProductDetailedScreen from './screens/ProductDetailedScreen'

import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom'

import CreateProductScreen from './screens/ProductCreateScreen'
import LoginScreen from './screens/LoginScreen'

export default class App extends React.Component {

    render() {
        return (
            <Router>
                <nav className="navbar navbar-expand-lg navbar-dark bg-dark static-top">
                    <div className="container">
                        <a className="navbar-brand" href="/">Lookum</a>
                        
                        <div className="collapse navbar-collapse" id="navbarResponsive">
                            <ul className="navbar-nav ml-auto">
                                <li className="nav-item active">
                                    <a className="nav-link" href="/">Home</a>
                                </li>
                                <li className="nav-item">
                                    <a className="nav-link" href="/createProduct">Create product</a>
                                </li>
                            </ul>

                        </div>
                        <ul className="navbar-nav ml-auto justify-content-end">
                            <li className="nav-item active">
                                <a className="nav-link" href="/login">Login</a>
                            </li>
                            <li className="nav-item">
                                <a className="nav-link" href="/register">Register</a>
                            </li>
                        </ul>
                    </div>
                </nav>

                <Switch>

                    <Route exact path="/">
                        <Products />
                    </Route>
                    <Route path="/createProduct" children={<CreateProductScreen />} />
                    <Route path="/product/:id" children={<ProductDetailedScreen />} />
                    <Route path="/login" children={<LoginScreen />} />
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