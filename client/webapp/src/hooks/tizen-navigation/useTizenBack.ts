import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const keyCodes = {
  BACK: 10009,
} as const;

function useTizenBack() {
  const navigate = useNavigate();

  useEffect(() => {
    const listener = (evt: KeyboardEvent) => {
      switch (evt.keyCode) {
        case keyCodes.BACK:
          navigate(-1);
          break;
        default:
          break;
      }
    };
    window.addEventListener("keydown", listener);

    return () => {
      window.removeEventListener("keydown", listener);
    };
  }, [navigate]);
}

export default useTizenBack;
