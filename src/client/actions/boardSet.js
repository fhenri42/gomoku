export default function boardSet(board) {

  return {
    type: 'client/boardSet',
    data: board,
  }
}
