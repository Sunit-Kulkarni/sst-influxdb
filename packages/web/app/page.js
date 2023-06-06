'use client';
import { useDisclosure, useInputState } from '@mantine/hooks';
import { Badge , Button, Image, Modal, Grid, TextInput } from "@mantine/core";

export default function Home() {
  const [opened, { open, close }] = useDisclosure(false);
  const [memberId, setMemberId] = useInputState('')
  const [payMessage, setPayMessage] = useInputState('')
  const fetchUser = async (id) => {
    const res = await fetch(`https://api-iot-dev.sunitkulkarni.com/member/${id}`);
    if (res.status == 200) {
      setPayMessage("Paid!")
      setTimeout(close, 1000)
     } else {
      setPayMessage("Invalid Member Id!")
      open()
     }
  }
  return (
    <div>
      <Modal opened={opened} onClose={close} title="Pay For Wine">
        <TextInput 
          label="Your Member ID" 
          description="Enter your country club membership number here"
          value={memberId}
          onChange={(event) => setMemberId(event.currentTarget.value)}
        />
        <br />
        <Button 
          color={'green'} 
          radius={'xl'} 
          size={'md'}
          onClick={() => fetchUser(memberId)}
        >
          Pay
        </Button>
        <h4>{payMessage}</h4>
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
        <Grid.Col span={6} align='center'>
          <Badge>VinoBarrel Estate</Badge>
          <Badge>raspberry</Badge>
          <Badge>cherry</Badge>
          <Badge>chocolate</Badge>
          <Badge>mixed-berry</Badge>
          <Badge>oak</Badge>
          <Badge>earthy</Badge>
          <br />
          <br />
          <br />
          <br />
          <Button onClick={open} radius={'xl'} size={'md'}>Single Pour $1.00</Button>
          <>&nbsp;&nbsp;&nbsp;&nbsp;</>
          <Button onClick={open} radius={'xl'} size={'md'}>Double Pour $2.00</Button>
        </Grid.Col>
      </Grid>
      <br />
      <br />
      <br />
      <br />
      <Grid align='center'>
        <Grid.Col span={6}>
          <Image
            src={"/VinoBarrelZinfandel.png"}
            maw={140}
            mx="auto"
            radius="sm"
          />
        </Grid.Col>
        <Grid.Col span={6} align='center'>
          <Badge>VinoBarrel Estate</Badge>
          <Badge>oak</Badge>
          <Badge>vanilla</Badge>
          <Badge>tobacco</Badge>
          <Badge>plum</Badge>
          <Badge>cherry</Badge>
          <Badge>dark fruit</Badge>
          <br />
          <br />
          <br />
          <br />
          <Button onClick={open} radius={'xl'} size={'md'}>Single Pour $1.00</Button>
          <>&nbsp;&nbsp;&nbsp;&nbsp;</>
          <Button onClick={open} radius={'xl'} size={'md'}>Double Pour $2.00</Button>
        </Grid.Col>
      </Grid>
    </div>
  )
}
