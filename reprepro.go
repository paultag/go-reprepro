/* {{{ Copyright (c) Paul R. Tagliamonte <paultag@debian.org>, 2015
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE. }}} */

package reprepro

import (
	"os/exec"
)

type Repo struct {
	Basedir string
}

func (repo *Repo) Command(args ...string) *exec.Cmd {
	return exec.Command("reprepro", append([]string{
		"--basedir", repo.Basedir,
	}, args...)...)
}

func (repo *Repo) ProcessIncoming(rule string) error {
	cmd := repo.Command("processincoming", rule)
	return cmd.Run()
}

func (repo *Repo) Check() error {
	cmd := repo.Command("check")
	return cmd.Run()
}

func (repo *Repo) CheckPool() error {
	cmd := repo.Command("checkpool")
	return cmd.Run()
}

func (repo *Repo) Include(suite string, changes string) error {
	cmd := repo.Command("include", suite, changes)
	return cmd.Run()
}

// Create a new reprepro.Repo object given a filesystem path to the Repo.
func NewRepo(path string) *Repo {
	return &Repo{Basedir: path}
}

// vim: foldmethod=marker
