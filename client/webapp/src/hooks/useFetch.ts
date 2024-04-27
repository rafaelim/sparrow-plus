import { useEffect, useState } from "react";

const useFetch = <T>(url: string) => {
  const [error, setError] = useState();
  const [data, setData] = useState<T>();

  useEffect(() => {
    fetch(url)
      .then((response) => response.json())
      .then((data) => {
        setData(data);
      })
      .catch((error) => {
        setError(error);
        return [];
      });
  }, [url]);

  return { data, error };
};

export default useFetch;
