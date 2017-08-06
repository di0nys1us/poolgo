'use strict';

import React from "react";
import { render } from "react-dom";

const Test = (props) => {
    return <div>
        <h1>
            {props.message}
        </h1>
    </div>
};

render(
    <Test message="Hello, world!" />,
    document.getElementById('root')
);