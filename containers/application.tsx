'use strict';

import * as React from "react";
import NavigationContainer from "./navigation"
import SessionList from "./session-list";

class ApplicationContainer extends React.Component<undefined, undefined> {
    render() {
        return <div>
            <NavigationContainer brand="PoolGo" />
            <SessionList />
        </div>
    }
}

export default ApplicationContainer