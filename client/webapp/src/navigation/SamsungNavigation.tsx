import { createContext, useEffect, useState } from "react";
import PositionHandler from "./PositionHandler";
import { useNavigate } from "react-router-dom";

const keyCodes = {
  LEFT: 37,
  UP: 38,
  RIGHT: 39,
  DOWN: 40,
  OK: 13,
  BACK: 10009,
} as const;

type NavigationContextValues = {
  position: { x: number; y: number };
  positionHandler: typeof PositionHandler;
};

export const NavigationContext = createContext<NavigationContextValues>({
  position: { x: 0, y: 0 },
  positionHandler: PositionHandler,
});

type SamsungNavigationProps = {
  children: React.ReactNode;
};

function SamsungNavigation({ children }: SamsungNavigationProps) {
  const navigate = useNavigate();
  const [position, setPosition] = useState({ x: 0, y: 0 });

  useEffect(() => {
    const listener = (evt: KeyboardEvent) => {
      switch (evt.keyCode) {
        case keyCodes.LEFT:
          setPosition((pos) => ({ x: Math.min(0, position.x - 1), y: pos.y }));
          break;
        case keyCodes.RIGHT:
          setPosition((pos) => ({ x: pos.x + 1, y: pos.y }));
          break;
        case keyCodes.UP:
          setPosition((pos) => ({ x: pos.x, y: Math.min(0, pos.y - 1) }));
          break;
        case keyCodes.DOWN:
          setPosition((pos) => ({
            x: pos.x,
            y: Math.min(PositionHandler.getLastY(), pos.y + 1),
          }));
          break;
        case keyCodes.OK:
          PositionHandler.triggerOpenEvent();
          break;
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
  });

  return (
    <NavigationContext.Provider
      value={{ position, positionHandler: PositionHandler }}
    >
      {children}
    </NavigationContext.Provider>
  );
}

export default SamsungNavigation;
