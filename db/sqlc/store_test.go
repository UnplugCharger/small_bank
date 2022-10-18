package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)




func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)

	account2 := createRandomAccount(t)


	n:= 5 
	amount := int64(10)

    errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++{
		go func(){
			result , err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID: account2.ID,
				Amount: amount,

			})
			results <- result
			errs <- err
		}()
	}
	// check results
	for i := 0; i < n; i++{
		result := <- results
		err := <- errs
		require.NoError(t, err)
		require.NotEmpty(t, result)
		
		transfer := result.Transfer
		require.NotZero(t, transfer.ID)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.CreatedAt)


		_, err = store.GetAccount(context.Background(), transfer.ID)
		require.NoError(t, err)

		//check entries 

		fromEntry := result.FromEntry
		require.NotZero(t, fromEntry.ID)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.CreatedAt)
		




	}
}