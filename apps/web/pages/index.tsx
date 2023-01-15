import React from 'react';
import Head from 'next/head';
import getConfig from 'next/config';
import { GetServerSideProps } from 'next';

const { publicRuntimeConfig } = getConfig();

type Biodata = {
  id: number;
  name: string;
};

export const getServerSideProps: GetServerSideProps = async () => {
  return { props: {} };
};

export default function Home() {
  const [biodatas, setBiodatas] = React.useState<Biodata[]>([]);

  React.useEffect(() => {
    fetch(publicRuntimeConfig.apiUrl)
      .then((res) => res.json())
      .then((res: Biodata[]) => setBiodatas(res));
  }, []);

  return (
    <div>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <p>List Biodatas :</p>
      <ol>
        {biodatas.map(({ id, name }) => (
          <li key={id}>{name}</li>
        ))}
      </ol>
    </div>
  );
}
