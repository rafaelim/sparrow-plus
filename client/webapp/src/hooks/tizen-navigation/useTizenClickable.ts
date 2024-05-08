import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const keyCodes = {
  OK: 13,
} as const;

type UseTizenClickableArgs = {
  onClick: () => void;
  isActive: boolean;
};

function useTizenClickable({ isActive, onClick }: UseTizenClickableArgs) {
  const navigate = useNavigate();

  useEffect(() => {
    const listener = (evt: KeyboardEvent) => {
      switch (evt.keyCode) {
        case keyCodes.OK:
          if (isActive) onClick();
          break;
        default:
          break;
      }
    };
    window.addEventListener("keydown", listener);

    return () => {
      window.removeEventListener("keydown", listener);
    };
  }, [onClick, navigate, isActive]);
}

export default useTizenClickable;
