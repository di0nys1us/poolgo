'use strict';

import * as React from "react";
import { render } from "react-dom";
import { Provider } from "react-redux";
import { combineReducers, createStore } from "redux";
import ApplicationContainer from "./containers/application"
import "./style/style"

const defaultReducer = (state = 0, action: any) => state;

const rootReducer = combineReducers({
    defaultReducer
});

const store = createStore(
    rootReducer
);

render(
    <Provider store={store}>
        <ApplicationContainer />
    </Provider>,
    document.getElementById('root')
);