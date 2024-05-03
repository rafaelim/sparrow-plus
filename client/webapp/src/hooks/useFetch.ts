/* eslint-disable @typescript-eslint/no-explicit-any */

import { useEffect, useState } from "react";

const useFetch = <T>(url: string) => {
  // const [error, setError] = useState();
  const a = useState<T>();
  console.log(a);
  useEffect(() => {
    // fetch(url)
    //   .then((response) => response.json())
    //   .then((data) => {
    //     setData(data);
    //   })
    //   .catch((error) => {
    //     setError(error);
    //     return [];
    //   });
  }, [url]);

  return {
    data: [
      { name: "test" },
      { name: "test" },
      { name: "test" },
      { name: "test" },
      { name: "test" },
      { name: "test" },
      { name: "test" },
      { name: "test" },
      { name: "test" },
      { name: "test" },
      { name: "test" },
      { name: "test" },
      { name: "test" },
      { name: "test" },
    ] as any,
    error: null,
  };
};

export default useFetch;
