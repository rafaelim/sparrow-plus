import MoviesList from "../../components/MoviesList";
import ShowList from "../../components/ShowList";
import WatchNext from "../../components/WatchNext";

const Home = () => {
  return (
    <>
      <WatchNext />
      <MoviesList />
      <ShowList />
    </>
  );
};

export default Home;
