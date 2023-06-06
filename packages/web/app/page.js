'use client';
import { useDisclosure } from '@mantine/hooks';
import { Input, Image, Modal, Grid, SimpleGrid, Box } from "@mantine/core";

export default function Home() {
  const [opened, { open, close }] = useDisclosure(false);
  return (
    <div>
      <Modal opened={opened} onClose={close} title="Enter Member ID">
        <Input />
      </Modal>
      <Grid align='center'>
        <Grid.Col span={6}>
          <Image
            src={"/VinoBarrelCabernet.png"}
            maw={140}
            mx="auto"
            radius="sm"
          />
        </Grid.Col>
        <Grid.Col span={6}>
          asdfasdfasdf
        </Grid.Col>
      </Grid>
      <Grid align='center'>
        <Grid.Col span={6}>
          <Image
            src={"/VinoBarrelZinfandel.png"}
            maw={140}
            mx="auto"
            radius="sm"
          />
        </Grid.Col>
        <Grid.Col span={6}>
          asdfasdfasdf
        </Grid.Col>
      </Grid>
    </div>
  )
}
