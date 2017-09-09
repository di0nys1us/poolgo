'use strict';

import * as React from "react";

interface NavigationContainerProps { brand: string }

class NavigationContainer extends React.Component<NavigationContainerProps, undefined> {
    render() {
        return <div className="navbar navbar-inverse bg-primary">
            <div className="container">
                <a className="navbar-brand" href="/">{this.props.brand}</a>
            </div>
        </div>
    }
}

export default NavigationContainer