import React from "react";
import {MainPageHeader} from '../mainPageHeader/index';
import "./index.css"

const MainPageContainer = () => {
    return (
        <div className="MainPageContainer-header">
            <MainPageHeader />
            <h2>Hello Dear User</h2>
        </div>
    )
};

export {MainPageContainer};