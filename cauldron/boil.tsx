interface Props {
  someProp: string;
}

const componentName: React.FC<Props> = (props) => {
  return (
    <>
      <h1>{props.someProp}</h1>
    </>
  );
};

export default componentName;
