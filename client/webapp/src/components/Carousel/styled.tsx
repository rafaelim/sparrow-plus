import styled from "styled-components";

export const Container = styled.div`
  margin: 10px;

  h1 {
    margin-top: 0;
  }
`;

export const ListContainer = styled.div`
  height: 300px;
  overflow-y: hidden;
  overflow-x: scroll;
  transition: 500ms;
  scroll-behavior: smooth;

  scrollbar-width: none;
  &::-webkit-scrollbar {
    display: none;
  }
`;

export const Body = styled.div`
  height: 100%;
  &:first-child {
    div {
      margin-left: 0;
    }
  }
  &:last-child {
    div {
      margin-right: 0;
    }
  }
`;

export const Title = styled.h1`
  text-transform: capitalize;
  font-family: sans-serif;
`;

// v.scroll-container {
//     height: 300px;
//     overflow-y: hidden;
//     overflow-x: scroll;
//     transition: 500ms;
//     scroll-behavior: smooth;

//     scrollbar-width: none;
//     &::-webkit-scrollbar {
//         display: none;
//     }
// }

// .card {
//     height: 100%;
//     &:first-child {
//         .box {
//             margin-left: 0;
//         }
//     }
//     &:last-child {
//         .box {
//             margin-right: 0;
//         }
//     }
// }
