import { useEffect } from "react";

const keyCodes = {
  LEFT: 37,
  UP: 38,
  RIGHT: 39,
  DOWN: 40,
  OK: 13,
  BACK: 10009,
} as const;

const SamsungNavigation = () => {
  useEffect(() => {
    const listener = (evt: KeyboardEvent) => {
      switch (evt.keyCode) {
        case keyCodes.LEFT:
          alert(keyCodes.LEFT);
          break;
        case keyCodes.RIGHT:
          alert(keyCodes.RIGHT);
          break;
        case keyCodes.UP:
          alert(keyCodes.UP);
          break;
        case keyCodes.DOWN:
          alert(keyCodes.DOWN);
          break;
        case keyCodes.OK:
          alert(keyCodes.OK);
          break;
        case keyCodes.BACK:
          alert(keyCodes.BACK);
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

  return <></>;
};

export default SamsungNavigation;
