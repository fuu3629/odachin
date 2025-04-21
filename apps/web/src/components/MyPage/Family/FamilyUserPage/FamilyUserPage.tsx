export interface FamilyUserPageProps {
  userId: string;
}

export function FamilyUserPage({ userId }: FamilyUserPageProps) {
  return (
    <>
      <div>{userId}</div>
    </>
  );
}
