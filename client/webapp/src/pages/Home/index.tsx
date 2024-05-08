import MoviesList from "../../components/MoviesList";
import ShowList from "../../components/ShowList";
import WatchNext from "../../components/WatchNext";
import SamsungNavigation from "../../navigation/SamsungNavigation";

const Home = () => {
  return (
    <SamsungNavigation>
      <WatchNext />
      <MoviesList />
      <ShowList />
    </SamsungNavigation>
  );
};

export default Home;
