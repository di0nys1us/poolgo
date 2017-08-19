'use strict';

import React from "react";
import { render } from "react-dom";
import { Provider } from "react-redux";
import { combineReducers, createStore } from "redux";
import SessionList from "./containers/session-list.js";

const defaultReducer = (state = 0, action) => state;

const rootReducer = combineReducers({
    defaultReducer
});

const store = createStore(
    rootReducer,
    window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
);

render(
    <Provider store={store}>
        <SessionList />
    </Provider>,
    document.getElementById('root')
);