import { FC, useEffect } from "react";
import { Icon } from "../components/ui/Icon";
import { PageSettings } from "./page";

class HomePageSettings extends PageSettings {}

const HomePage: FC = () => {
  const settings = new HomePageSettings("Home", true, true, false);
  useEffect(() => {
    settings.setup();
  }, []);
  return (
    <div>
      <h1>Home Page</h1>
      <Icon name="browser" />
    </div>
  );
};

export default HomePage;
