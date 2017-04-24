export default function activation(data) {
  console.log('lelel');
  return {
    type: 'server/activation',
    data,
  }
}
