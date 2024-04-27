import { createGlobalStyle } from "styled-components";

const GlobalStyle = createGlobalStyle`
    * {
        box-sizing: border-box;
    }
    html, 
    body, 
    #root {
        width: 100%;
        height: 100%;
        margin: 0;
        padding: 0;
        background-color: #D6FFF6;
    }
`;

export default GlobalStyle;
