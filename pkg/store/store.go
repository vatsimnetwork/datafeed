package store

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

type Store struct {
	dataServerURL  string
	transceiverURL string
	atisClientsURL string

	DataServerJSON  string
	TransceiverJSON string
	ATISClientsJSON string
}

func New(dataServerURL, transceiverURL, atisClientsURL string) *Store {
	return &Store{
		dataServerURL:  dataServerURL,
		transceiverURL: transceiverURL,
		atisClientsURL: atisClientsURL,
	}
}

func (s *Store) Run() {
	fmt.Println("Updating store...")
	if err := s.Update(); err != nil {
		fmt.Printf("Error updating store: %s\n", err)
	}

	ticker := time.NewTicker(15 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println("Updating store...")
			if err := s.Update(); err != nil {
				fmt.Printf("Error updating store: %s\n", err)
			}
		}
	}
}

func (s *Store) Update() error {
	g := errgroup.Group{}

	g.Go(s.fetchDataServerJSON)
	g.Go(s.fetchTransceiverJSON)
	g.Go(s.fetchATISClientsJSON)

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func (s *Store) fetchDataServerJSON() error {
	body, err := fetchURL(s.dataServerURL)
	if err != nil {
		return err
	}

	s.DataServerJSON = body

	return nil
}

func (s *Store) fetchTransceiverJSON() error {
	body, err := fetchURL(s.transceiverURL)
	if err != nil {
		return err
	}

	s.TransceiverJSON = body

	return nil
}

func (s *Store) fetchATISClientsJSON() error {
	body, err := fetchURL(s.atisClientsURL)
	if err != nil {
		return err
	}

	s.ATISClientsJSON = body

	return nil
}

func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
