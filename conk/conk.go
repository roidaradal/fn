// Package conk contains concurrency-related functions.
package conk

import "golang.org/x/sync/errgroup"

// Run concurrent task for each item,
// Task takes item as input and outputs (data, error).
// onReceive processes each output data.
func Tasks[T, D any](items []T, task func(T) (D, error), onReceive func(D)) error {
	// Run goroutine for each item
	var eg errgroup.Group
	dataCh := make(chan D, len(items))
	for _, item := range items {
		eg.Go(func() error {
			result, err := task(item)
			if err != nil {
				return err
			}
			dataCh <- result
			return nil
		})
	}

	// Wait for errgroup and close data channel
	var finalErr error
	go func() {
		finalErr = eg.Wait()
		close(dataCh)
	}()

	// Receive from data channel
	for result := range dataCh {
		onReceive(result)
	}
	return finalErr
}
